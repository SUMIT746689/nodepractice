import { configureStore } from '@reduxjs/toolkit'
import users from './features/users'
import { userApi } from './services/user'
import { setupListeners } from '@reduxjs/toolkit/dist/query'
import { authApi } from './services/auth';

export const store = configureStore({
  reducer: {
    user: users,
    [userApi.reducerPath]:userApi.reducer,
    [authApi.reducerPath]:authApi.reducer,
  },
  devTools: true,
  middleware: (getDefaultMiddleware)=>getDefaultMiddleware()
  .concat(userApi.middleware)
  .concat(authApi.middleware)
});


setupListeners(store.dispatch);

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch