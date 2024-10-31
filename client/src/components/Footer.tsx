import { Box, Stack, Text, Container, Divider } from "@chakra-ui/react";

const Footer = () =>
{
    return (
        <Container width="100%" maxW="100%">
            <Box
                color="white"
                px={ 4 }
            >
                <Stack direction={ { base: 'column', md: 'row' } }
                    justifyContent="space-between"
                    spacing={ 5 } py={ 4 }
                >
                    <Text fontSize="2xs" color="gray.600" textAlign="center">
                        © { new Date().getFullYear() } by buwu
                    </Text>
                </Stack>
                <Divider borderColor="gray.600" />
                <Stack direction={ { base: 'column', md: 'row' } }
                    justifyContent="space-between"
                    spacing={ 5 } py={ 4 }
                >
                    <Text fontSize="2xs" color="gray.600" textAlign="center">
                        ♥
                    </Text>
                </Stack>
            </Box>
        </Container>
    );
};

export default Footer;