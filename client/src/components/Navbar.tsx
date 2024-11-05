import {Box, Container, useColorMode, Flex, Text, Button, Link, Tabs, TabList, Tab, Spinner, Drawer, DrawerOverlay, DrawerContent, DrawerCloseButton, DrawerHeader, DrawerBody, VStack, IconButton} from "@chakra-ui/react";
import {useMutation} from "@tanstack/react-query";
import {useState} from "react";
import {BiLogOutCircle} from "react-icons/bi";
import {IoMoon} from "react-icons/io5";
import {LuSun} from "react-icons/lu";
import {GiHamburgerMenu} from "react-icons/gi";
import {useNavigate} from "react-router-dom";

export default function Navbar()
{
    const {colorMode, toggleColorMode} = useColorMode();
    const [newTodo, setNewTodo] = useState("");
    const navigate = useNavigate();
    const [isSidebarOpen, setSidebarOpen] = useState(false);

    const {mutate: logout, isPending: isLogout} = useMutation({
        mutationKey: ["createTodo"],
        mutationFn: async (e: React.FormEvent) =>
        {
            e.preventDefault();
            localStorage.removeItem("jwt");
            try
            {
                const response = await fetch("http://localhost:5000/api/logout", {
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
        <Container>
            <Box bgGradient="linear(to-l, teal.800, purple.800)" px={4} borderRadius="5">
                <Flex h={20} alignItems="center" justifyContent="space-between">
                    {/* LEFT */}
                    <Flex alignItems="center" gap={5}>
                        <IconButton
                            icon={<GiHamburgerMenu />}
                            aria-label="Open sidebar"
                            onClick={() => setSidebarOpen(true)}
                        />
                        <Text fontSize="lg" fontWeight={500}>
                            go-n-note
                        </Text>
                    </Flex>

                    {/* RIGHT */}
                    <Flex alignItems="center" gap={5}>
                        <Button onClick={toggleColorMode}>
                            {colorMode === "light" ? <IoMoon /> : <LuSun size={20} />}
                        </Button>
                        <Button onClick={logout}>
                            {isLogout ? <Spinner size="xs" /> : <BiLogOutCircle size={20} />}
                        </Button>
                    </Flex>
                </Flex>
            </Box>

            {/* Tabs */}
            <Box px={4} alignItems="center" py={2}>
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

            {/* Sidebar Drawer */}
            <Drawer isOpen={isSidebarOpen} placement="left" onClose={() => setSidebarOpen(false)}>
                <DrawerOverlay />
                <DrawerContent bgGradient="linear(to-l, teal.800, purple.800)">
                    <DrawerCloseButton color="white" />
                    <DrawerHeader color="white">Menu</DrawerHeader>
                    <DrawerBody>
                        <VStack align="start" spacing={4}>
                            <Link href="#" _hover={{color: "teal.300"}}>
                                <Text color="gray.300" fontWeight="thin">Home</Text>
                            </Link>
                            <Link href="#" _hover={{color: "teal.300"}}>
                                <Text color="gray.300" fontWeight="thin">Dashboard</Text>
                            </Link>
                            <Link href="#" _hover={{color: "teal.300"}}>
                                <Text color="gray.300" fontWeight="thin">Settings</Text>
                            </Link>
                            <Link href="#" _hover={{color: "teal.300"}}>
                                <Text color="gray.300" fontWeight="thin">Profile</Text>
                            </Link>
                        </VStack>
                    </DrawerBody>
                </DrawerContent>
            </Drawer>
        </Container>
    );
}
