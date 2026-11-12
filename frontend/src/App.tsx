
import { Box, Button, Container } from "@chakra-ui/react"
import { useState } from "react"


function App() {
  const [count, setCount] = useState(0)

  return (
    <Container>
      <Box>
        Count : {count}
      </Box>
      

      <Button onClick={()=>setCount(count+1)}>Submit</Button>
      
    </Container>
  )
}

export default App
