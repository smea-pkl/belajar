import { PasswordInput } from "@/components/ui/password-input";
import { toaster } from "@/components/ui/toaster"
import { Button, Card, Center, Container, Input } from "@chakra-ui/react";
import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient } from "@connectrpc/connect";
import { UserService } from "@/gen/pkl/users/v1/user_service_pb"
import { useState } from "react";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});


const userServiceClient = createClient(UserService, transport)


export default function() {

    const [name, setName] = useState()
    const [username, setUsername] = useState()
    const [password, setPassword] = useState()



    const register = () => {

        const payload = {
            name: name,
            password: password,
            username: username
        }

        userServiceClient.
            register(payload).
            then(()=>{
                toaster.create({
                    title: "Register",
                    description: "Register Berhasil",
                    type: "success",
                    
                })
            }).
            catch(()=>{
                toaster.create({
                    title: "Register",
                    description: "Register Gagal",
                    type: "error",
                    
                })
            })

        
        
    }

    return (
        <Container>

              <Center mt="16">
                <Card.Root w="40%">
                  <Card.Body gap="2">
                    
                    <Card.Title mt="2">Register</Card.Title>
                    <Card.Description>
                      Silahkan Login Untuk Berjudi.
                    </Card.Description>
                    <Input placeholder="Name" value={name} onChange={e=>setName(e.target.value as any)}/>
                    <Input placeholder="Username" value={username} onChange={e=>setUsername(e.target.value as any)} />
                    <PasswordInput placeholder="Password" value={password} onChange={e=>setPassword(e.target.value as any)} />
                  </Card.Body>
                  <Card.Footer justifyContent="flex-end">
                    <Button variant="outline" size="sm" onClick={register}>Register</Button>
                  </Card.Footer>
                </Card.Root>
              </Center>
              
            </Container>
        
    )
}