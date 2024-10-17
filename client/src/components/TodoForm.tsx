import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";

const TodoForm = () =>
{
    const [newTodo, setNewTodo] = useState("");

    const queryClient = useQueryClient();

    const { mutate: createTodo, isPending: isCreating } = useMutation({
        mutationKey: ["createTodo"],
        mutationFn: async (e: React.FormEvent) =>
        {
            e.preventDefault()
            try
            {
                const response = await fetch("/api/todos", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({ body: newTodo })
                })
                const data = await response.json()
                console.log(data)
                if (!response.ok)
                {
                    throw new Error(data.message)
                }
                setNewTodo("")
                return data
            }
            catch (error: any)
            {
                throw new Error(error.message)
            }
        },
        onSuccess: () =>
        {
            queryClient.invalidateQueries({ queryKey: ["todos"] })
        },
        onError: (error: any) =>
        {
            alert(error.message)
        }
    })
    return (
        <form onSubmit={ (e) => { e.preventDefault(); createTodo(e); } }>
            <Flex >
                <Input
                    type='text'
                    value={ newTodo }
                    onChange={ (e) => setNewTodo(e.target.value) }
                    ref={ (input) => input && input.focus() }
                />
                <Button
                    mx={ 2 }
                    type='submit'

                >
                    { isCreating ? <Spinner size={ "xs" } /> : <IoMdAdd size={ 30 } /> }
                </Button>
            </Flex>
        </form>
    );
};
export default TodoForm; 