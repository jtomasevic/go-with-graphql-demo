"use client"

import Router, { useRouter } from "next/router";
import { useGetShipsQuery } from "../../api/gql";
import Link from "next/link";
import { ApolloClient, ApolloProvider, NormalizedCacheObject } from '@apollo/client';
import appoloClient from "@/api";

export default function Index() {
    return (
        <ApolloProvider client={appoloClient}>
            <ShipList />
        </ApolloProvider>
    )
}

export function ShipList() {
    const { loading, error, data } = useGetShipsQuery();
    if (loading) {
        return <p>Loading...</p>;
    }
    if (error) {
        return <p>Error : {error.message}</p>;
    }
    return data!.ships.map(({ id, name }) => (
        <div key={id}>
            <h3><Link  href={'/ships/' + id}>{name}</Link></h3>
        </div>
    ));
}