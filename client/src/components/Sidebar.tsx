// Sidebar.tsx
import
{
    Box,
    VStack,
    Text,
    Drawer,
    DrawerOverlay,
    DrawerContent,
    DrawerCloseButton,
    DrawerHeader,
    DrawerBody,
    useBreakpointValue,
    IconButton,
    Button,
    Divider,
    Icon,
} from "@chakra-ui/react";
import {GiHamburgerMenu} from "react-icons/gi";
import {FaHome, FaChartLine, FaCog, FaUser} from "react-icons/fa"; // Import icons
import {useState} from "react";

export default function Sidebar()
{
    const [isDrawerOpen, setDrawerOpen] = useState(false);
    const isDrawer = useBreakpointValue({base: true, md: false});

    // Define items with labels and icons
    const menuItems = [
        {label: 'Home', icon: FaHome},
        {label: 'Dashboard', icon: FaChartLine},
        {label: 'Settings', icon: FaCog},
        {label: 'Profile', icon: FaUser},
    ];

    return (
        <>
            {isDrawer ? (
                <>
                    <IconButton
                        icon={<GiHamburgerMenu />}
                        aria-label="Open sidebar"
                        onClick={() => setDrawerOpen(true)}
                        m={4}
                    />
                    <Drawer isOpen={isDrawerOpen} placement="left" onClose={() => setDrawerOpen(false)}>
                        <DrawerOverlay />
                        <DrawerContent background={"#841e7452"}>
                            <DrawerCloseButton color="white" />
                            <DrawerHeader color="white">Menu</DrawerHeader>
                            <DrawerBody>
                                <VStack align="start" spacing={0}>
                                    {menuItems.map((item) => (
                                        <Box key={item.label} width="full">
                                            <Button
                                                variant="outline"
                                                width="full"
                                                justifyContent="start"
                                                px={4}
                                                py={2}
                                                borderRadius="md"
                                                borderColor="transparent"
                                                _hover={{
                                                    color: "white",
                                                    borderColor: "#00ffff",
                                                    backgroundColor: "whiteAlpha.100",
                                                    border: "1px solid",
                                                    transition: "all 0.3s ease"
                                                }}
                                                color="gray.300"
                                                fontWeight="thin"
                                            >
                                                <Icon as={item.icon} mr={3} />
                                                {item.label}
                                            </Button>
                                        </Box>
                                    ))}
                                </VStack>
                            </DrawerBody>
                        </DrawerContent>
                    </Drawer>
                </>
            ) : (
                <Box
                    background={"#841e7452"}
                    w="250px"
                    h="100vh"
                    p={4}
                    color="white"
                >
                    <Text fontSize="xl" fontWeight="semibold" mb={4} textAlign="left">GONOTE</Text>
                    <Divider></Divider>
                    <VStack align="start" spacing={0}>
                        {menuItems.map((item, index) => (
                            <Box key={item.label} width="full">
                                <Button
                                    variant="outline"
                                    width="full"
                                    justifyContent="start"
                                    px={4}
                                    py={2}
                                    borderRadius="md"
                                    borderColor="transparent"
                                    _hover={{
                                        color: "white",
                                        backgroundColor: "whiteAlpha.200",
                                        transition: "all 0.3s ease"
                                    }}
                                    fontWeight="normal"
                                >
                                    <Icon as={item.icon} mr={3} /> {/* Add icon */}
                                    {item.label}
                                </Button>
                            </Box>
                        ))}
                    </VStack>
                </Box>
            )}
        </>
    );
}
