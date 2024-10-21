import { Box, Container, useColorModeValue, useColorMode, Flex, Text, Button, Divider, Stack, Link, Tabs, TabList, Tab, TabPanels, TabPanel } from "@chakra-ui/react";
import { IoMoon } from "react-icons/io5";
import { LuSun } from "react-icons/lu";
import gopher from "../../public/gopher.jpg";

export default function Navbar()
{
    const { colorMode, toggleColorMode } = useColorMode();
    return (
        <Container maxW={ "900px" }>
            <Box
                px={ 4 } alignItems={ "center" } py={ 2 }
            >
                <Tabs variant={ "line" } align="center" >
                    <TabList>
                        <Tab>
                            <Link href="#" _hover={ { color: 'white' } }>
                                <Text fontWeight={ "thin" } fontSize={ "14" } color={ "gray.400" }>
                                    Home
                                </Text>
                            </Link>
                        </Tab>
                        <Tab>
                            <Link href="#" _hover={ { color: 'white' } }>
                                <Text fontWeight={ "thin" } fontSize={ "14" } color={ "gray.300" }>
                                    About
                                </Text>
                            </Link>
                        </Tab>
                        <Tab>
                            <Link href="#" _hover={ { color: 'white' } }>
                                <Text fontWeight={ "thin" } fontSize={ "14" } color={ "gray.300" }>
                                    Contact
                                </Text>
                            </Link>
                        </Tab>
                    </TabList>
                </Tabs>

            </Box>
            <Box bg={ useColorModeValue("purple.100", "purple.800") }
                px={ 4 }
                borderRadius={ "5" }
            >
                <Flex h={ 20 } alignItems={ "right" } justifyContent={ "space-between" } >
                    {/* LEFT */ }
                    <Flex>
                        <img src={ gopher } alt="logo" height={ 5 } width={ 100 }></img>
                    </Flex>
                    {/* RIGHT */ }
                    <Flex alignItems={ "center" } gap={ 5 }>
                        <Text fontSize="lg" fontWeight={ 500 }>
                            Daily Tasks
                        </Text>
                        <Button onClick={ toggleColorMode }>
                            { colorMode === "light" ? <IoMoon /> : <LuSun size={ 20 } /> }
                        </Button>
                    </Flex>
                </Flex>
            </Box>
        </Container >
    )
}
