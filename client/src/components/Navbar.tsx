import { Box, Container, useColorModeValue, useColorMode, Flex, Text, Button } from "@chakra-ui/react";
import { IoMoon } from "react-icons/io5";
import { LuSun } from "react-icons/lu";
import gopher from "../../public/gopher.jpg";

export default function Navbar()
{
    const { colorMode, toggleColorMode } = useColorMode();
    return (
        <Container maxW={ "900px" }>
            <Box bg={ useColorModeValue("purple.100", "purple.800") } px={ 4 } my={ 4 } borderRadius={ "5" } >
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
