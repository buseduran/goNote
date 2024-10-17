import { Badge, Box, Flex, Spinner, Text } from '@chakra-ui/react'
import { FaCheckCircle } from 'react-icons/fa'
import { MdDelete } from 'react-icons/md'
import { Todo } from './TodoList'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { BASE_URL } from '../App'

const TodoItem = ({ todo }: { todo: Todo }) =>
{
    const queryClient = useQueryClient()
    const { mutate: updateTodo, isPending: isUpdating } = useMutation({
        mutationKey: ["updateTodo"],
        mutationFn: async () =>
        {
            if (todo.completed)
                return alert("Task already completed")

            try
            {
                const response = await fetch(BASE_URL + `/todos/${ todo._id }`, {
                    method: "PATCH"
                })
                const data = await response.json()
                if (!response.ok)
                {
                    throw new Error(data.message || "Something went wrong")
                }
                return data
            }
            catch (error)
            {
                console.log(error)
            }
        },
        onSuccess: () =>
        {
            queryClient.invalidateQueries({ queryKey: ["todos"] })
        }
    })
    const { mutate: deleteTodo, isPending: isDeleting } = useMutation({
        mutationKey: ["deleteTodo"],
        mutationFn: async () =>
        {
            try
            {
                const response = await fetch(BASE_URL + `/todos/${ todo._id }`, {
                    method: "DELETE"
                })
                const data = await response.json()
                if (!response.ok)
                {
                    throw new Error(data.message || "Something went wrong")
                }
                return data
            }
            catch (error)
            {
                console.log(error)
            }
        },
        onSuccess: () =>
        {
            queryClient.invalidateQueries({ queryKey: ["todos"] })
        }
    })
    return (
        <Flex alignItems={ "center" } gap={ 3 }>
            <Flex
                flex={ 1 }
                alignItems={ "center" }
                border={ "1px" }
                borderColor={ "gray.700" }
                padding={ 3 }
                borderRadius={ "lg" }
                justifyContent={ "space-between" }
            >
                <Text color={ todo.completed ? "green.200" : "yellow.200" }
                    textDecoration={ todo.completed ? "line-through" : "none" }>
                    { todo.body }
                </Text>
                { todo.completed && (
                    <Badge ml='1' colorScheme='green'>
                        Done
                    </Badge>
                ) }
                { !todo.completed && (
                    <Badge ml='1' colorScheme='yellow'>
                        In Progress
                    </Badge>
                ) }
            </Flex>
            <Flex alignItems={ "center" } gap={ 2.5 }>
                <Box color={ "green.500" } cursor={ "pointer" } onClick={ () => updateTodo() }>
                    { !isUpdating && <FaCheckCircle size={ 20 } /> }
                    { isUpdating && <Spinner size={ "sm" } /> }
                </Box>
                <Box color={ "red.500" } cursor={ "pointer" } onClick={ () => deleteTodo() }>
                    { isDeleting && <Spinner size={ "sm" } /> }
                    { !isDeleting && <MdDelete size={ 20 } /> }
                </Box>
            </Flex>
        </Flex>
    )
}


export default TodoItem