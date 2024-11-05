import {Stack, Text} from "@chakra-ui/react"
import TodoItem from "./TodoItem"
import {useQuery} from "@tanstack/react-query"

export type Todo = {
    id: number
    body: string
    completed: boolean
}

const TodoList = () =>
{
    //fetch datas
    const {data: todos, isLoading, error} = useQuery<Todo[]>({
        queryKey: ["todos"],
        queryFn: async () =>
        {
            try
            {
                const response = await fetch("http://localhost:5000/api/todos", {
                    method: "GET",
                    headers: {
                        "Authorization": `Bearer ${ localStorage.getItem("jwt") }`,
                        "Content-Type": "application/json"
                    },
                    credentials: "include"
                })
                const data = await response.json()
                if (!response.ok)
                {
                    throw new Error(data.message || "Something went wrong")
                }
                return data || []
            }
            catch (error)
            {
                console.log(error)
            }
        }
    })

    if (isLoading)
    {
        return <Text>Loading...</Text>
    }
    if (error)
    {
        return <Text>Error: {error.message}</Text>
    }

    return (
        <>
            <Text textTransform={"uppercase"}
                fontSize={"xl"}
                fontWeight={"semibold"}
                textAlign={"center"}
                my={2}
                bgGradient='linear(to-l, purple.500, #00ffff)'
                bgClip={"text"}>
                Tasks
            </Text >
            {!isLoading && todos?.length === 0 && (
                <Stack>
                    <Text textAlign={"center"}
                        fontWeight={"thin"}
                        textDecoration={"ButtonFace"}>
                        No tasks for today 🙌
                    </Text>
                </Stack>
            )}
            <Stack gap={3}>
                {todos?.map((todo) => (
                    <TodoItem key={todo.id} todo={todo} />
                ))}
            </Stack>
        </>
    )
}

export default TodoList