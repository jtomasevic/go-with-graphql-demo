import appoloClient from "@/api";
import { useGetPiratesQuery } from "@/api/gql";
import { ApolloProvider } from "@apollo/client";


export default function Pirates() {
  return (
    <ApolloProvider client={appoloClient}>
      <PiratesList />
    </ApolloProvider>
  )
}


export function PiratesList() {
    const { loading, error, data } = useGetPiratesQuery();
    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error : {error.message}</p>;
    return data!.pirates.map(({ id, name }) => (
      <div key={id}>
        <h3>{name}</h3>
      </div>
    ));
  }