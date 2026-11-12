import React from 'react';
import { Container } from "@chakra-ui/react"
import NavigationBar from "./NavigationBar"

function Layout({children }: React.PropsWithChildren){
    return (
        <Container>
            <NavigationBar />
            {children}
        </Container>
    )
}


export default Layout

