import { API_KEY } from '@/secret'
import { AuthLogIn, AuthUser } from '@/types/auth'
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'


export const authApi = createApi({
  reducerPath: 'authApi',
  baseQuery: fetchBaseQuery({ baseUrl: API_KEY }),
  endpoints: (builder) => ({
    loginUser: builder.mutation<AuthLogIn, string>({
      query: (body) => ({
        url: 'login',
        method: 'POST',
        body,
      })
    }),
    authUser: builder.query({
      query: () => ({
        url:'me',
        credentials:'include'
      }),
    }),
  }),

})

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const { useLoginUserMutation, useAuthUserQuery } = authApi