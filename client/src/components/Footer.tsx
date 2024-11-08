import {Box, Stack, Text, Container, Divider} from "@chakra-ui/react";

const Footer = () =>
{
    return (
        <Container width="100%" maxW="100%" centerContent>
            <Box
                color="white"
                px={4}
                alignItems="center"
                justifyContent="center"
                textAlign="center"
                width={"100%"}
                position={"fixed"}
                bottom={0}
                zIndex={1}
            >
                {/* <Stack
                    direction={{ base: 'column', md: 'row' }}
                    justifyContent="center"
                    alignItems="center"
                    spacing={5}
                    py={4}
                >
                    <Text fontSize="2xs" color="gray.600">
                        ♥
                    </Text>
                </Stack> */}
                <Divider borderColor="gray.600" w="80%" mx="auto" />

                <Stack
                    direction={{base: 'column', md: 'row'}}
                    justifyContent="center"
                    alignItems="center"
                    spacing={5}
                    py={1}
                >
                    <Text fontSize="2xs" color="gray.600">
                        © {new Date().getFullYear()} by buwu  ♥
                    </Text>
                </Stack>
            </Box>
        </Container>
    );
};

export default Footer;
