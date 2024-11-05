// App.tsx
import {Container, Stack} from "@chakra-ui/react";
import Navbar from "./components/Navbar";
import TodoForm from "./components/todo-components/TodoForm";
import TodoList from "./components/todo-components/TodoList";
import Footer from "./components/Footer";
import Sidebar from "./components/Sidebar";

export const BASE_URL = import.meta.env.MODE === "/api";

function App()
{
  return (
    <Stack minH="100vh" direction="row">
      <Sidebar />
      <Stack flex="1">
        <Navbar />
        <Container flex="1">
          <TodoForm />
          <TodoList />
        </Container>
        <Footer />
      </Stack>
    </Stack>
  );
}

export default App;
