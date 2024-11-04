import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { ChakraProvider } from '@chakra-ui/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import theme from './chakra/theme.ts'
import Login from './Login.tsx'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Signup from './components/login-components/SignupForm.tsx'

const queryClient = new QueryClient()


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={ queryClient }>
      <ChakraProvider theme={ theme }>
        <BrowserRouter>
          <Routes>
            <Route path='/' element={ <App /> } />
            <Route path='/login' element={ <Login /> } />
            <Route path='/register' element={ <Signup /> } />
          </Routes>
        </BrowserRouter>
      </ChakraProvider>
    </QueryClientProvider>
  </StrictMode>
)
