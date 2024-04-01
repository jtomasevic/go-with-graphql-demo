"use client"

import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client';
import appoloClient from '@/api';
import Link from 'next/link';



export default function Home() {
  return (
    <>
        <h1>Pirates of Caribbean</h1>
        <div>
            <Link href="/ships">Ships</Link>
        </div>
    </>
  )
}