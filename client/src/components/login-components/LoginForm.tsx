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
import {useMutation, useQueryClient} from "@tanstack/react-query";
import {useState} from "react";
import {FaUserAlt, FaLock} from "react-icons/fa";
import {useNavigate} from "react-router-dom";

const CFaUserAlt = chakra(FaUserAlt);
const CFaLock = chakra(FaLock);


const Login = () =>
{
    const [showPassword, setShowPassword] = useState(false)

    const handleShowClick = () => setShowPassword(!showPassword)

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const queryClient = useQueryClient()
    const navigate = useNavigate()

    const {mutate: login, isPending: isCreating} = useMutation({
        mutationKey: ["loginUser"],
        mutationFn: async (e: React.FormEvent) =>
        {
            e.preventDefault()
            try
            {
                const response = await fetch("/api/login", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({username, password})
                })
                const data = await response.json()
                console.log(data)
                if (!response.ok)
                {
                    throw new Error(data.message)
                }
                return data
            }
            catch (error: any)
            {
                throw new Error(error.message)
            }
        },
        onSuccess: (data) =>
        {
            localStorage.setItem("jwt", data.token);
            queryClient.invalidateQueries({queryKey: ["user"]})
            navigate("/")
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
                <Avatar bg="#841e7452" />
                <Heading
                    bgGradient='linear(to-l, purple.500, #00ffff)'
                    bgClip={"text"}
                    mt={2}
                    fontWeight={"semibold"}
                    fontSize={"2xl"}
                >Welcome</Heading>
                <Box minW={{base: "90%", md: "468px"}}>
                    <form onSubmit={(e) => {e.preventDefault; login(e);}}>
                        <Stack
                            spacing={4}
                            p="1rem"
                            boxShadow="md"
                        >
                            <FormControl>
                                <InputGroup>
                                    <InputLeftElement
                                        pointerEvents="none"
                                        children={<CFaUserAlt color="gray.300" />}
                                    />
                                    <Input
                                        type="text"
                                        value={username}
                                        placeholder="username"
                                        onChange={(e) => setUsername(e.target.value)} />
                                </InputGroup>
                            </FormControl>
                            <FormControl>
                                <InputGroup>
                                    <InputLeftElement
                                        pointerEvents="none"
                                        color="gray.300"
                                        children={<CFaLock color="gray.300" />}
                                    />
                                    <Input
                                        type={showPassword ? "text" : "password"}
                                        placeholder="password"
                                        value={password}
                                        onChange={(e) => setPassword(e.target.value)}
                                    />
                                    <InputRightElement width="4.5rem">
                                        <Button h="1.75rem" size="sm" onClick={handleShowClick}>
                                            {showPassword ? "Hide" : "Show"}
                                        </Button>
                                    </InputRightElement>
                                </InputGroup>
                                <FormHelperText textAlign="right">
                                    <Link>forgot password?</Link>
                                </FormHelperText>
                            </FormControl>
                            <Button
                                borderRadius={2}
                                type="submit"
                                variant="solid"
                                background={"#841e7452"}
                                width="full"
                                isDisabled={isCreating}
                            >
                                {isCreating ? <Spinner size={"xs"} /> : <Text>Login</Text>}
                            </Button>
                        </Stack>
                    </form>
                </Box>
            </Stack>
            <Box>
                New to us?{" "}
                <Link background={"#00ffff"} bgClip="text" href="/register">
                    Sign Up
                </Link>
            </Box>
        </Flex >
    );
}

export default Login