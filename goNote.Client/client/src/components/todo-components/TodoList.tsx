import {Stack, Text} from "@chakra-ui/react"
import TodoItem from "./TodoItem"
import {useQuery} from "@tanstack/react-query"
import {useState} from "react"

export type Todo = {
    id: number
    body: string
    completed: boolean
}

const TodoList = () =>
{
    const [redirecting, setRedirecting] = useState(false);

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
                    credentials: "include" //if using cookies
                })

                if (response.status === 401)
                {
                    setRedirecting(true)
                    //window.location.href = BASE_URL + "/login"
                    //return new Promise(() => {})
                }

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
                return Promise.reject(error)
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

    if (redirecting) return null; // Prevent rendering while redirecting
    if (isLoading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

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