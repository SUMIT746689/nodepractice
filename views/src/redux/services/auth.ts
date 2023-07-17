import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'


export const authApi = createApi({
  reducerPath: 'authApi',
  baseQuery: fetchBaseQuery({ baseUrl: 'http://127.0.0.1:8080/api/v1/' }),
  endpoints: (builder) => ({
    loginUser: builder.mutation<any, string>({
      query: (body) => ({
        url: 'login',
        method: 'POST',
        body,
      })
    }),
  }),

})

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const { useLoginUserMutation } = authApi