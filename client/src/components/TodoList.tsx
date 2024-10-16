import { Stack, Text } from "@chakra-ui/react"
import { useState } from "react"
import TodoItem from "./TodoItem"

const TodoList = () =>
{
    const [isLoading, setIsLoading] = useState(true)
    const todos = [
        {
            id: 1,
            body: "Task 1",
            completed: false
        },
        {
            id: 2,
            body: "Task 2",
            completed: true
        }
    ]
    return (
        <>
            <Text textTransform={ "uppercase" }
                fontSize={ "xl" }
                fontWeight={ "semibold" }
                textAlign={ "center" }
                my={ 2 }
                bgGradient='linear(to-l, purple.600, #00ffff)'
                bgClip={ "text" }>
                Today's Tasks
            </Text >
            { !isLoading && todos?.length === 0 && (
                <Stack>
                    <Text>
                        No tasks for today
                    </Text>
                    <img src="/gopher.jpg" width={ 100 } height={ 100 }></img>
                </Stack>
            )
            }
            <Stack gap={ 3 }>
                { todos.map((todo) => (
                    <TodoItem key={ todo.id } todo={ todo } />
                )) }
            </Stack>
        </>
    )
}

export default TodoList