// Navbar.tsx
import
{
    Box,
    Flex,
    Text,
    Button,
    Spinner,
    useColorMode,
    Link,
    Tabs,
    TabList,
    Tab
} from "@chakra-ui/react";
import {useMutation} from "@tanstack/react-query";
import {useState} from "react";
import {BiLogOutCircle} from "react-icons/bi";
import {IoMoon} from "react-icons/io5";
import {LuSun} from "react-icons/lu";
import {useNavigate} from "react-router-dom";

export default function Navbar()
{
    const {colorMode, toggleColorMode} = useColorMode();
    const navigate = useNavigate();
    const [newTodo, setNewTodo] = useState("");

    const {mutate: logout, isPending: isLogout} = useMutation({
        mutationKey: ["createTodo"],
        mutationFn: async (e: React.FormEvent) =>
        {
            e.preventDefault();
            localStorage.removeItem("jwt");
            try
            {
                const response = await fetch("/logout", {
                    method: "POST",
                    headers: {
                        "Authorization": `Bearer ${ localStorage.getItem("jwt") }`,
                        "Content-Type": "application/json"
                    },
                    credentials: "include",
                    body: JSON.stringify({body: newTodo})
                });

                const data = await response.json();
                if (!response.ok) throw new Error(data.message);
                setNewTodo("");
                return data;
            } catch (error: any)
            {
                throw new Error(error.message);
            }
        },
        onSuccess: () => navigate("/login"),
        onError: (error: any) => alert(error.message)
    });

    return (
        <Box px={4} py={4} borderRadius="5">
            <Flex h={10} alignItems="center" justifyContent="right">
                <Flex alignItems="center" gap={5}>
                    <Button onClick={toggleColorMode}>
                        {colorMode === "light" ? <IoMoon /> : <LuSun size={20} />}
                    </Button>
                    <Button onClick={logout}>
                        {isLogout ? <Spinner size="xs" /> : <BiLogOutCircle size={20} />}
                    </Button>
                </Flex>
            </Flex>
            {/* Navigation Tabs */}
            <Tabs variant="line" align="center" colorScheme="blue">
                <TabList>
                    <Tab>
                        <Link href="#" _hover={{color: "white"}}>
                            <Text fontWeight="thin" fontSize="14" color="gray.300">Daily</Text>
                        </Link>
                    </Tab>
                    <Tab>
                        <Link href="#" _hover={{color: "white"}}>
                            <Text fontWeight="thin" fontSize="14" color="gray.300">Weekly</Text>
                        </Link>
                    </Tab>
                    <Tab>
                        <Link href="#" _hover={{color: "white"}}>
                            <Text fontWeight="thin" fontSize="14" color="gray.300">Monthly</Text>
                        </Link>
                    </Tab>
                </TabList>
            </Tabs>
        </Box>
    );
}
