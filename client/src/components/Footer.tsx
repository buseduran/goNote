import { Box, Stack, Text, Container } from "@chakra-ui/react"

const Footer = () =>
{
    return (

        <Container maxW={ "900px" } >
            <Box
                color={ "white" }
                marginTop={ 6 }
                px={ 4 }
            >
                <Stack direction={ { base: 'column', md: 'row' } }
                    justifyContent={ "space-between" }
                    spacing={ 5 } py={ 4 }
                >
                    <Text fontSize={ "2xs" } color={ "gray.600" } justifyItems={ "center" } >
                        © { new Date().getFullYear() } by buwu
                    </Text>
                </Stack>
            </Box>
        </Container>
    )
}

export default Footer