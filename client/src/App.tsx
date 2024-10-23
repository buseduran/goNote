import { Container, Stack } from '@chakra-ui/react'
import Navbar from './components/Navbar'
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'
import Footer from './components/Footer'

export const BASE_URL = import.meta.env.MODE === "/api"

function App()
{
  return (
    <Stack minH="100vh" position={ "sticky" }>
      <Navbar />
      <Container flex="1">
        <TodoForm />
        <TodoList />
      </Container>
      <Footer />
    </Stack>
  )
}


export default App
