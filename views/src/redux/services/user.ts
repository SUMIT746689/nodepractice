import { API_KEY } from '@/secret';
import { User } from '@/types/users'
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
// import axios from "axios"

// const axiosBaseQuery = ({ baseUrl } = { baseUrl: '' }) =>
//   async ({ url, method, data, params }) => {
//     try {
//       const result = await axios({ url: baseUrl + url, method, data, params, withCredentials:true })
//       console.log({result})
//       return { data: result.data }
//     } catch (axiosError) {
//       let err = axiosError
//       console.error({err})
//       return {
//         error: {
//           status: err.response?.status,
//           data: err.response?.data || err.message,
//         },
//       }
//     }
//   }

export const userApi = createApi({
  reducerPath: 'userApi',
  tagTypes: ['Users'],
  baseQuery: fetchBaseQuery({
    baseUrl: API_KEY,
    credentials: 'include'
  }),
  // refetchOnReconnect: true,
  endpoints: (builder) => ({
    getAllUsers: builder.query({
      query: () => ("/users"),
        // providesTags: ['Users'],
        // transformResponse: (response, meta, arg) => {
        //   console.log({ a: response })
        //   return 'hi'
        // },
        // // Pick out errors and prevent nested properties in a hook or selector
        // transformErrorResponse: (response, meta, arg) => {
        //   console.log({ response })
        //   return "errrr..."
        //   // return response.status
        // },
      // }),
      providesTags: ['Users'],
    }),
    postUser: builder.mutation<User, string>({
      query: (body) => ({
        url: '/users',
        method: 'POST',
        body
      }),
      // invalidatesTags: ['Users'],
      // transformResponse: (result: { data: { users: any } }) =>result.data.users)
    }),

    updateUser: builder.mutation<User, string>({
      query: (body) => ({
        url: '/users',
        method: 'PUT',
        body
      })
    }),
  }),

})

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const { useGetAllUsersQuery, usePostUserMutation, useUpdateUserMutation } = userApi;