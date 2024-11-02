import { Container, Stack } from '@chakra-ui/react'
import Footer from './components/Footer'
import LoginForm from './components/login-components/LoginForm'

export const BASE_URL = import.meta.env.MODE === "/api"

function Login() {
    return (
        <Stack>
            <Container flex="1">
                <LoginForm />
                <Footer />
            </Container>
        </Stack>
    )
}

export default Login
