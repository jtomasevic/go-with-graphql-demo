import appoloClient from "@/api";
import { useGetPiratesQuery, useGetShipQuery } from "@/api/gql";
import { ApolloProvider } from "@apollo/client";
import { useRouter } from 'next/router'

export default function Ship() {
    return (
        <ApolloProvider client={appoloClient}>
            <ShipDetails/>
        </ApolloProvider>
    )
}


export function ShipDetails() {
    const router = useRouter()
    console.log(router.query.ship);
    const { loading, error, data } = useGetShipQuery({ variables:{ id: router.query.ship}});
    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error : {error.message}</p>;
    const pirates = data!.ship.crew.pirates.map(({ id, name }) => (
        <div key={id}>
          <h5>{name}</h5>
        </div>
    ))
    return (
        <>
        <h2>Ship: {data?.ship.name}</h2>
        <h3>Crew: {data?.ship.crew.name}</h3>
        <h4>Pirates: </h4>
        {pirates}
        </>
    );
  }