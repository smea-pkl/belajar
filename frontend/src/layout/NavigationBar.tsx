import { Box, Button, Flex, HStack} from "@chakra-ui/react"
import { Link } from "react-router"

function NavigationBar(){
    var isLogin = false


    var loginedMenu: React.JSX.Element

    if(isLogin ){
      loginedMenu = <Button size="sm" variant="ghost">Profile</Button>
    } else {
      loginedMenu = (
        <Link to={"/login"}>
          <Button size="sm" variant="ghost">Login</Button>
        </Link>
      )
    }

    return (
        <Box>
          <Flex mx="auto" mt="4" align="center" justify="flex-end">
            <HStack>
                <Link to={"/"}>
                  <Button size="sm">Home</Button>
                </Link>
                
                <Button size="sm" variant="ghost">Catur</Button>
                <Button size="sm" variant="ghost">Judol</Button>
                {loginedMenu}
                
                
                
            </HStack>
          </Flex>
        </Box>
    )
}

export default NavigationBar