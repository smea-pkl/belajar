import Layout from "@/layout/Layout"
import { Button, HStack, Stack } from "@chakra-ui/react";
import { useState } from "react";

const products = [
  { title: 'Cabbage', id: 1 },
  { title: 'Garlic', id: 2 },
  { title: 'Apple', id: 3 },
];

function Home(){

    const [count, setCount] = useState(0)

    const tambah = () => {
        setCount(count + 1)
    }

   
    return (
        <Layout>
            {count}
            <Button onClick={tambah}>Tambah</Button>
            <HStack>
                {products.map(product =>
                    <Button key={product.id}>
                        {product.title}
                    </Button>
                )}
            </HStack>

            

        </Layout>
    )
}

export default Home