import { Alert, AlertDescription, AlertIcon, AlertTitle, Badge, Box, Flex, Spinner, Text } from '@chakra-ui/react'
import { FaCheckCircle, FaRegCircle } from 'react-icons/fa'
import { MdDelete } from 'react-icons/md'
import { Todo } from './TodoList'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useState } from 'react'

const TodoItem = ({ todo }: { todo: Todo }) =>
{
    const queryClient = useQueryClient()
    const [showAlert, setShowAlert] = useState(false)
    const [alertMessage, setAlertMessage] = useState('')
    const [alertStatus, setAlertStatus] = useState<'error' | 'success'>('success')

    //UPDATE
    const { mutate: updateTodo, isPending: isUpdating } = useMutation({
        mutationKey: ["updateTodo"],
        mutationFn: async () =>
        {
            try
            {
                const response = await fetch(`http://localhost:5000/api/todos/${ todo._id }`, {
                    method: "PATCH",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({ completed: !todo.completed })
                })
                if (!response.ok)
                {
                    const errorMessage = await response.text();
                    setAlertStatus('error');
                    setAlertMessage(errorMessage || "Something went wrong");
                    setShowAlert(true);
                    return;
                }

                const data = await response.json();
                return data;
            } catch (error)
            {
                console.log(error)
            }
        },
        onSuccess: () =>
        {
            queryClient.invalidateQueries({ queryKey: ["todos"] })
        }
    })

    //DELETE
    const { mutate: deleteTodo, isPending: isDeleting } = useMutation({
        mutationKey: ["deleteTodo"],
        mutationFn: async () =>
        {
            try
            {
                const response = await fetch(`http://localhost:5000/api/todos/${ todo._id }`, {
                    method: "DELETE"
                })
                const data = await response.json()
                if (!response.ok)
                {
                    setAlertStatus('error')
                    setAlertMessage(data.message || "Something went wrong")
                    setShowAlert(true)
                    return;
                }
                return data
            } catch (error)
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
        <>
            { showAlert && (
                <Alert status={ alertStatus }>
                    <AlertIcon />
                    <AlertTitle mr={ 2 }>{ alertStatus === 'error' ? 'Error' : 'Success' }</AlertTitle>
                    <AlertDescription>{ alertMessage }</AlertDescription>
                </Alert>
            ) }

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
                        { !isUpdating && (todo.completed ? <FaCheckCircle size={ 20 } /> : <FaRegCircle size={ 18 }></FaRegCircle>) }
                        { isUpdating && <Spinner size={ "sm" } /> }
                    </Box>
                    <Box color={ "red.500" } cursor={ "pointer" } onClick={ () => deleteTodo() }>
                        { isDeleting && <Spinner size={ "sm" } /> }
                        { !isDeleting && <MdDelete size={ 20 } /> }
                    </Box>
                </Flex>
            </Flex>
        </>
    )
}

export default TodoItem
