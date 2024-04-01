import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
import { graphql } from 'msw'
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  UUID: { input: any; output: any; }
};

export type Crew = {
  __typename?: 'Crew';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  pirates: Array<Pirate>;
};

export type Mutation = {
  __typename?: 'Mutation';
  createCrew: Crew;
  createPirate: Pirate;
  createShip: Ship;
};


export type MutationCreateCrewArgs = {
  input: UpsertCrew;
};


export type MutationCreatePirateArgs = {
  input: UpsertPirate;
};


export type MutationCreateShipArgs = {
  input: UpsertShip;
};

export type Pirate = {
  __typename?: 'Pirate';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type Query = {
  __typename?: 'Query';
  crews: Array<Crew>;
  pirates: Array<Pirate>;
  ship: Ship;
  ships: Array<Ship>;
};


export type QueryShipArgs = {
  id?: InputMaybe<Scalars['UUID']['input']>;
};

export type Ship = {
  __typename?: 'Ship';
  crew: Crew;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type UpsertCrew = {
  id?: InputMaybe<Scalars['ID']['input']>;
  name: Scalars['String']['input'];
  shipId?: InputMaybe<Scalars['UUID']['input']>;
};

export type UpsertPirate = {
  crewId?: InputMaybe<Scalars['UUID']['input']>;
  id?: InputMaybe<Scalars['ID']['input']>;
  name: Scalars['String']['input'];
};

export type UpsertShip = {
  id?: InputMaybe<Scalars['ID']['input']>;
  name: Scalars['String']['input'];
};

export type GetPiratesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetPiratesQuery = { __typename?: 'Query', pirates: Array<{ __typename?: 'Pirate', id: string, name: string }> };

export type GetCrewsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetCrewsQuery = { __typename?: 'Query', crews: Array<{ __typename?: 'Crew', id: string, name: string, pirates: Array<{ __typename?: 'Pirate', id: string, name: string }> }> };

export type GetShipsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetShipsQuery = { __typename?: 'Query', ships: Array<{ __typename?: 'Ship', id: string, name: string }> };

export type GetShipQueryVariables = Exact<{
  id: Scalars['UUID']['input'];
}>;


export type GetShipQuery = { __typename?: 'Query', ship: { __typename?: 'Ship', id: string, name: string, crew: { __typename?: 'Crew', id: string, name: string, pirates: Array<{ __typename?: 'Pirate', id: string, name: string }> } } };


export const GetPiratesDocument = gql`
    query GetPirates {
  pirates {
    id
    name
  }
}
    `;

/**
 * __useGetPiratesQuery__
 *
 * To run a query within a React component, call `useGetPiratesQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetPiratesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetPiratesQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetPiratesQuery(baseOptions?: Apollo.QueryHookOptions<GetPiratesQuery, GetPiratesQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetPiratesQuery, GetPiratesQueryVariables>(GetPiratesDocument, options);
      }
export function useGetPiratesLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetPiratesQuery, GetPiratesQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetPiratesQuery, GetPiratesQueryVariables>(GetPiratesDocument, options);
        }
export function useGetPiratesSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<GetPiratesQuery, GetPiratesQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<GetPiratesQuery, GetPiratesQueryVariables>(GetPiratesDocument, options);
        }
export type GetPiratesQueryHookResult = ReturnType<typeof useGetPiratesQuery>;
export type GetPiratesLazyQueryHookResult = ReturnType<typeof useGetPiratesLazyQuery>;
export type GetPiratesSuspenseQueryHookResult = ReturnType<typeof useGetPiratesSuspenseQuery>;
export type GetPiratesQueryResult = Apollo.QueryResult<GetPiratesQuery, GetPiratesQueryVariables>;
export const GetCrewsDocument = gql`
    query GetCrews {
  crews {
    id
    name
    pirates {
      id
      name
    }
  }
}
    `;

/**
 * __useGetCrewsQuery__
 *
 * To run a query within a React component, call `useGetCrewsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetCrewsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetCrewsQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetCrewsQuery(baseOptions?: Apollo.QueryHookOptions<GetCrewsQuery, GetCrewsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetCrewsQuery, GetCrewsQueryVariables>(GetCrewsDocument, options);
      }
export function useGetCrewsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetCrewsQuery, GetCrewsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetCrewsQuery, GetCrewsQueryVariables>(GetCrewsDocument, options);
        }
export function useGetCrewsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<GetCrewsQuery, GetCrewsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<GetCrewsQuery, GetCrewsQueryVariables>(GetCrewsDocument, options);
        }
export type GetCrewsQueryHookResult = ReturnType<typeof useGetCrewsQuery>;
export type GetCrewsLazyQueryHookResult = ReturnType<typeof useGetCrewsLazyQuery>;
export type GetCrewsSuspenseQueryHookResult = ReturnType<typeof useGetCrewsSuspenseQuery>;
export type GetCrewsQueryResult = Apollo.QueryResult<GetCrewsQuery, GetCrewsQueryVariables>;
export const GetShipsDocument = gql`
    query GetShips {
  ships {
    id
    name
  }
}
    `;

/**
 * __useGetShipsQuery__
 *
 * To run a query within a React component, call `useGetShipsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetShipsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetShipsQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetShipsQuery(baseOptions?: Apollo.QueryHookOptions<GetShipsQuery, GetShipsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetShipsQuery, GetShipsQueryVariables>(GetShipsDocument, options);
      }
export function useGetShipsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetShipsQuery, GetShipsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetShipsQuery, GetShipsQueryVariables>(GetShipsDocument, options);
        }
export function useGetShipsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<GetShipsQuery, GetShipsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<GetShipsQuery, GetShipsQueryVariables>(GetShipsDocument, options);
        }
export type GetShipsQueryHookResult = ReturnType<typeof useGetShipsQuery>;
export type GetShipsLazyQueryHookResult = ReturnType<typeof useGetShipsLazyQuery>;
export type GetShipsSuspenseQueryHookResult = ReturnType<typeof useGetShipsSuspenseQuery>;
export type GetShipsQueryResult = Apollo.QueryResult<GetShipsQuery, GetShipsQueryVariables>;
export const GetShipDocument = gql`
    query GetShip($id: UUID!) {
  ship(id: $id) {
    id
    name
    crew {
      id
      name
      pirates {
        id
        name
      }
    }
  }
}
    `;

/**
 * __useGetShipQuery__
 *
 * To run a query within a React component, call `useGetShipQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetShipQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetShipQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useGetShipQuery(baseOptions: Apollo.QueryHookOptions<GetShipQuery, GetShipQueryVariables> & ({ variables: GetShipQueryVariables; skip?: boolean; } | { skip: boolean; }) ) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetShipQuery, GetShipQueryVariables>(GetShipDocument, options);
      }
export function useGetShipLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetShipQuery, GetShipQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetShipQuery, GetShipQueryVariables>(GetShipDocument, options);
        }
export function useGetShipSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<GetShipQuery, GetShipQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<GetShipQuery, GetShipQueryVariables>(GetShipDocument, options);
        }
export type GetShipQueryHookResult = ReturnType<typeof useGetShipQuery>;
export type GetShipLazyQueryHookResult = ReturnType<typeof useGetShipLazyQuery>;
export type GetShipSuspenseQueryHookResult = ReturnType<typeof useGetShipSuspenseQuery>;
export type GetShipQueryResult = Apollo.QueryResult<GetShipQuery, GetShipQueryVariables>;

/**
 * @param resolver a function that accepts a captured request and may return a mocked response.
 * @see https://mswjs.io/docs/basics/response-resolver
 * @example
 * mockGetPiratesQuery((req, res, ctx) => {
 *   return res(
 *     ctx.data({ pirates })
 *   )
 * })
 */
export const mockGetPiratesQuery = (resolver: Parameters<typeof graphql.query<GetPiratesQuery, GetPiratesQueryVariables>>[1]) =>
  graphql.query<GetPiratesQuery, GetPiratesQueryVariables>(
    'GetPirates',
    resolver
  )

/**
 * @param resolver a function that accepts a captured request and may return a mocked response.
 * @see https://mswjs.io/docs/basics/response-resolver
 * @example
 * mockGetCrewsQuery((req, res, ctx) => {
 *   return res(
 *     ctx.data({ crews })
 *   )
 * })
 */
export const mockGetCrewsQuery = (resolver: Parameters<typeof graphql.query<GetCrewsQuery, GetCrewsQueryVariables>>[1]) =>
  graphql.query<GetCrewsQuery, GetCrewsQueryVariables>(
    'GetCrews',
    resolver
  )

/**
 * @param resolver a function that accepts a captured request and may return a mocked response.
 * @see https://mswjs.io/docs/basics/response-resolver
 * @example
 * mockGetShipsQuery((req, res, ctx) => {
 *   return res(
 *     ctx.data({ ships })
 *   )
 * })
 */
export const mockGetShipsQuery = (resolver: Parameters<typeof graphql.query<GetShipsQuery, GetShipsQueryVariables>>[1]) =>
  graphql.query<GetShipsQuery, GetShipsQueryVariables>(
    'GetShips',
    resolver
  )

/**
 * @param resolver a function that accepts a captured request and may return a mocked response.
 * @see https://mswjs.io/docs/basics/response-resolver
 * @example
 * mockGetShipQuery((req, res, ctx) => {
 *   const { id } = req.variables;
 *   return res(
 *     ctx.data({ ship })
 *   )
 * })
 */
export const mockGetShipQuery = (resolver: Parameters<typeof graphql.query<GetShipQuery, GetShipQueryVariables>>[1]) =>
  graphql.query<GetShipQuery, GetShipQueryVariables>(
    'GetShip',
    resolver
  )
