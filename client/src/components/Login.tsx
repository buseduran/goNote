import
{
    Avatar,
    Box,
    Button,
    chakra,
    Flex,
    FormControl,
    FormHelperText,
    Heading,
    Input,
    InputGroup,
    InputLeftElement,
    InputRightElement,
    Link,
    Spinner,
    Stack,
    Text
} from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { FaUserAlt, FaLock } from "react-icons/fa";

const CFaUserAlt = chakra(FaUserAlt);
const CFaLock = chakra(FaLock);


const Login = () =>
{
    const [showPassword, setShowPassword] = useState(false);

    const handleShowClick = () => setShowPassword(!showPassword);

    const [newTodo, setNewTodo] = useState("");

    const queryClient = useQueryClient();

    const { mutate: createTodo, isPending: isCreating } = useMutation({
        mutationKey: ["createTodo"],
        mutationFn: async (e: React.FormEvent) =>
        {
            e.preventDefault()
            try
            {
                const response = await fetch("http://localhost:5000/api/todos", {
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
        <Flex
            flexDirection="column"
            width="100wh"
            height="100vh"
            backgroundColor="gray.800"
            justifyContent="center"
            alignItems="center"
        >
            <Stack
                flexDir="column"
                mb="2"
                justifyContent="center"
                alignItems="center"
            >
                <Avatar bg="teal.500" />
                <Heading color="teal.400">Welcome</Heading>
                <Box minW={ { base: "90%", md: "468px" } }>
                    <form onSubmit={ (e) => { e.preventDefault; loginUser(e); } }>
                        <Stack
                            spacing={ 4 }
                            p="1rem"
                            backgroundColor="whiteAlpha.100"
                            boxShadow="md"
                        >
                            <FormControl>
                                <InputGroup>
                                    <InputLeftElement
                                        pointerEvents="none"
                                        children={ <CFaUserAlt color="gray.300" /> }
                                    />
                                    <Input type="username" placeholder="username" />
                                </InputGroup>
                            </FormControl>
                            <FormControl>
                                <InputGroup>
                                    <InputLeftElement
                                        pointerEvents="none"
                                        color="gray.300"
                                        children={ <CFaLock color="gray.300" /> }
                                    />
                                    <Input
                                        type={ showPassword ? "text" : "password" }
                                        placeholder="password"
                                    />
                                    <InputRightElement width="4.5rem">
                                        <Button h="1.75rem" size="sm" onClick={ handleShowClick }>
                                            { showPassword ? "Hide" : "Show" }
                                        </Button>
                                    </InputRightElement>
                                </InputGroup>
                                <FormHelperText textAlign="right">
                                    <Link>forgot password?</Link>
                                </FormHelperText>
                            </FormControl>
                            <Button
                                borderRadius={ 2 }
                                type="submit"
                                variant="solid"
                                colorScheme="teal"
                                width="full"
                            >
                                { isCreating ? <Spinner size={ "xs" } /> : <Text>Login</Text> }
                            </Button>
                        </Stack>
                    </form>
                </Box>
            </Stack>
            <Box>
                New to us?{ " " }
                <Link color="teal.500" href="#">
                    Sign Up
                </Link>
            </Box>
        </Flex >
    );
}

export default Login