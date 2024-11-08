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
import {BsTextRight} from "react-icons/bs";
import {FaLock} from "react-icons/fa";
import {FaUserAstronaut} from "react-icons/fa6";
import {useNavigate} from "react-router-dom";
import {BASE_URL} from "../../App";

const CFaLock = chakra(FaLock);

const Signup = () =>
{
    const [showPassword, setShowPassword] = useState(false)

    const handleShowClick = () => setShowPassword(!showPassword)

    const [firstname, setFirstname] = useState("")
    const [lastname, setLastname] = useState("")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const queryClient = useQueryClient()
    const navigate = useNavigate()

    const {mutate: signup, isPending: isCreating} = useMutation({
        mutationKey: ["signupUser"],
        mutationFn: async () =>
        {
            try
            {
                const response = await fetch(BASE_URL + "/register", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({firstname: firstname, lastname: lastname, username: username, password: password})
                });
                const data = await response.json();
                if (!response.ok)
                {
                    throw new Error(data.message);
                }
                return data;
            } catch (error: any)
            {
                throw new Error(error.message);
            }
        },
        onSuccess: () =>
        {
            queryClient.invalidateQueries({queryKey: ["user"]});
            navigate("/login");
        },
        onError: (error: any) =>
        {
            alert(error.message);
        }
    });


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
                    <form onSubmit={(e) => {e.preventDefault(); signup();}}>
                        <Stack
                            spacing={4}
                            p="1rem"
                            boxShadow="md"
                        >
                            <FormControl>
                                <InputGroup>
                                    <InputLeftElement
                                        pointerEvents="none"
                                        children={<BsTextRight color="gray.300" />}
                                    />
                                    <Input
                                        type="text"
                                        value={firstname}
                                        placeholder="firstname"
                                        onChange={(e) => setFirstname(e.target.value)} />
                                </InputGroup>
                            </FormControl>
                            <FormControl>
                                <InputGroup>
                                    <InputLeftElement
                                        pointerEvents="none"
                                        children={<BsTextRight color="gray.300" />}
                                    />
                                    <Input
                                        type="text"
                                        value={lastname}
                                        placeholder="lastname"
                                        onChange={(e) => setLastname(e.target.value)} />
                                </InputGroup>
                            </FormControl>
                            <FormControl>
                                <InputGroup>
                                    <InputLeftElement
                                        pointerEvents="none"
                                        children={<FaUserAstronaut color="gray.300" />}
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
                                {isCreating ? <Spinner size={"xs"} /> : <Text>Signup</Text>}
                            </Button>
                        </Stack>
                    </form>
                </Box>
            </Stack>
            <Box>
                Have an account?{" "}
                <Link background={"#00ffff"} bgClip="text" href="/login">
                    Log In
                </Link>
            </Box>
        </Flex >
    );
}

export default Signup