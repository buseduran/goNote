import { Badge, Box, Flex, Text } from '@chakra-ui/react'
import { FaCheckCircle } from 'react-icons/fa'
import { MdDelete } from 'react-icons/md'

const TodoItem = ({ todo }: { todo: any }) =>
{
    return (
        <Flex alignItems={ "center" } gap={ 3 }>
            <Flex
                flex={ 1 }
                alignItems={ "center" }
                border={ "1px" }
                borderColor={ "gray.700" }
                padding={ 3 }
                borderRadius={ "lg" }
                justifyContent={ "space-between" }
            >
                <Text color={ todo.completed ? "green.200" : "yellow.200" }
                    textDecoration={ todo.completed ? "line-through" : "none" }>
                    { todo.body }
                </Text>
                { todo.completed && (
                    <Badge ml='1' colorScheme='green'>
                        Done
                    </Badge>
                ) }
                { !todo.completed && (
                    <Badge ml='1' colorScheme='yellow'>
                        In Progress
                    </Badge>
                ) }
            </Flex>
            <Flex alignItems={ "center" } gap={ 2.5 }>
                <Box color={ "green.500" } cursor={ "pointer" }>
                    <FaCheckCircle size={ 20 } />
                </Box>
                <Box color={ "red.500" } cursor={ "pointer" }>
                    <MdDelete size={ 20 } />
                </Box>
            </Flex>
        </Flex>
    )
}


export default TodoItem