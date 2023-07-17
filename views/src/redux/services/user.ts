import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'


export const pokemonApi = createApi({
  reducerPath: 'pokemonApi',
  baseQuery: fetchBaseQuery({ baseUrl: 'https://pokeapi.co/api/v2/' }),
  endpoints: (builder) => ({
    getPokemonByName: builder.query<any, string>({
      query: (form) => ({
        url: 'login',
        method: 'POST',
        body: form
      }),
    }),
  }),

})

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const { useGetPokemonByNameQuery } = pokemonApi