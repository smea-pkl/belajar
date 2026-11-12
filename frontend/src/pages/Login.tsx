
import { PasswordInput, PasswordStrengthMeter } from "@/components/ui/password-input"
import { Center, Container, Input } from "@chakra-ui/react"
import { Button, Card } from "@chakra-ui/react"


function Login() {


  return (
    <Container>
      <Center mt="16">
        <Card.Root w="40%">
          <Card.Body gap="2">
            
            <Card.Title mt="2">Login</Card.Title>
            <Card.Description>
              Silahkan Login Untuk Berjudi.
            </Card.Description>

            <Input placeholder="Username" />
            <PasswordInput placeholder="Password"/>
            <PasswordStrengthMeter value={2}/>
          </Card.Body>
          <Card.Footer justifyContent="flex-end">
            <Button size="sm">Login</Button>
            <Button variant="outline" size="sm">Register</Button>
            
          </Card.Footer>
        </Card.Root>
      </Center>
      
    </Container>
  )
}

export default Login
