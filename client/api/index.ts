import { ApolloClient, InMemoryCache } from "@apollo/client";

const appoloClient = new ApolloClient({
    uri: 'http://127.0.0.1:8080/query',
    cache: new InMemoryCache(),
  });

  export default appoloClient;