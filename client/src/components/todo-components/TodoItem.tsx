import {Alert, AlertDescription, AlertIcon, AlertTitle, Badge, Box, Flex, Input, Spinner} from '@chakra-ui/react'
import {FaCheckCircle, FaRegCircle} from 'react-icons/fa'
import {MdDelete} from 'react-icons/md'
import {useMutation, useQueryClient} from '@tanstack/react-query'
import {useState} from 'react'
import {Todo} from './TodoList'

const TodoItem = ({todo}: {todo: Todo}) =>
{
    const queryClient = useQueryClient()
    const [showAlert, setShowAlert] = useState(false)
    const [alertMessage, setAlertMessage] = useState('')
    const [alertStatus, setAlertStatus] = useState<'error' | 'success'>('success')
    const [description, setDescription] = useState(todo.body)

    //UPDATE BODY
    const {mutate: updateTodoBody} = useMutation({
        mutationKey: ["updateTodoBody"],
        mutationFn: async () =>
        {
            try
            {
                const response = await fetch(`/api/todos/${ todo.id }`, {
                    method: "PATCH",
                    headers: {
                        "Authorization": `Bearer ${ localStorage.getItem("jwt") }`,
                        "Content-Type": "application/json"
                    },
                    credentials: "include",
                    body: JSON.stringify({completed: todo.completed, body: description})
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
            queryClient.invalidateQueries({queryKey: ["todos"]})
        }
    })

    //UPDATE CHECK
    const {mutate: updateTodoToggle, isPending: isUpdatingToggle} = useMutation({
        mutationKey: ["updateTodoToggle"],
        mutationFn: async () =>
        {
            try
            {
                console.log("datalarr " + todo)
                const response = await fetch(`/api/todos/${ todo.id }`, {
                    method: "PATCH",
                    headers: {
                        "Authorization": `Bearer ${ localStorage.getItem("jwt") }`,
                        "Content-Type": "application/json"
                    },
                    credentials: "include",
                    body: JSON.stringify({completed: !todo.completed, body: description})
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
            queryClient.invalidateQueries({queryKey: ["todos"]})
        }
    })


    //DELETE
    const {mutate: deleteTodo, isPending: isDeleting} = useMutation({
        mutationKey: ["deleteTodo"],
        mutationFn: async () =>
        {
            try
            {
                const response = await fetch(`/api/todos/${ todo.id }`, {
                    headers: {
                        "Authorization": `Bearer ${ localStorage.getItem("jwt") }`,
                        "Content-Type": "application/json"
                    },
                    credentials: "include",
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
            queryClient.invalidateQueries({queryKey: ["todos"]})
        }
    })

    return (
        <>
            {showAlert && (
                <Alert status={alertStatus}>
                    <AlertIcon />
                    <AlertTitle mr={2}>{alertStatus === 'error' ? 'Error' : 'Success'}</AlertTitle>
                    <AlertDescription>{alertMessage}</AlertDescription>
                </Alert>
            )}

            <Flex alignItems={"center"} gap={3}>
                <Flex
                    flex={1}
                    alignItems={"center"}
                    border={"1px"}
                    borderColor={"gray.700"}
                    padding={3}
                    borderRadius={"lg"}
                    justifyContent={"space-between"}
                >
                    <Input
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        onBlur={() => updateTodoBody()}
                        size={"sm"}
                        placeholder='Edit Todo'
                        minWidth={"200px"}
                        maxWidth={"300px"}
                        border={'none'}
                        textDecoration={todo.completed ? "line-through" : "none"}
                    >
                    </Input>

                    {todo.completed && (
                        <Badge ml='1' colorScheme='green'>
                            Done
                        </Badge>
                    )}
                    {!todo.completed && (
                        <Badge ml='1' colorScheme='yellow'>
                            In Progress
                        </Badge>
                    )}
                </Flex>
                <Flex alignItems={"center"} gap={2.5}>
                    <Box color={"green.500"} cursor={"pointer"} onClick={() => updateTodoToggle()}>
                        {!isUpdatingToggle && (todo.completed ? <FaCheckCircle size={20} /> : <FaRegCircle size={18}></FaRegCircle>)}
                        {isUpdatingToggle && <Spinner size={"sm"} />}
                    </Box>
                    <Box color={"red.500"} cursor={"pointer"} onClick={() => deleteTodo()}>
                        {isDeleting && <Spinner size={"sm"} />}
                        {!isDeleting && <MdDelete size={20} />}
                    </Box>
                </Flex>
            </Flex>
        </>
    )
}

export default TodoItem
