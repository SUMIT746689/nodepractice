import { API_KEY } from '@/secret';
import { GetAllUsersInterface, UpdateUser, User } from '@/types/users'
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
    getAllUsers: builder.query<User[],void>({
      query: () => ("/users"),
      transformResponse: (response: GetAllUsersInterface) => response.users,
      providesTags: [{ type: "Users", id: "LIST" }],
    }),
    postUser: builder.mutation<User, string>({
      query: (body) => ({
        url: '/users',
        method: 'POST',
        body,
        // headers: { "Content-Type": "text/plain" }
      }),
      invalidatesTags: [{ type: "Users", id: "LIST" }],
      // transformResponse: (result: { data: { users: any } }) =>result.data.users)
    }),

    updateUser: builder.mutation<User, UpdateUser>({
      query: ({ user_id, body }) => ({
        url: `/users/${JSON.stringify(user_id)}`,
        method: 'PATCH',
        body,
        responseHandler:(response)=> response.text(),
      }),
      invalidatesTags: [{ type: "Users", id: "LIST" }],
    }),
    deleteUser: builder.mutation<User, number>({
      query: (user_id) => ({
        url: `/users/${JSON.stringify(user_id)}`,
        method: 'DELETE',
        responseHandler: (response) =>  response.text(),
      }),
      invalidatesTags: ['Users']
    }),
  }),

})

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const { useGetAllUsersQuery, usePostUserMutation, useUpdateUserMutation, useDeleteUserMutation } = userApi;