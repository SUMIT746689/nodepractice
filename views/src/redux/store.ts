import { configureStore } from '@reduxjs/toolkit'
import users from './features/users'
import { userApi } from './services/user'
import { setupListeners } from '@reduxjs/toolkit/dist/query'
import { authApi } from './services/auth';
import { roleApi } from './services/role';
import { companyApi } from './services/company';

export const store = configureStore({
  reducer: {
    user: users,
    [userApi.reducerPath]: userApi.reducer,
    [authApi.reducerPath]: authApi.reducer,
    [roleApi.reducerPath]: roleApi.reducer,
    [companyApi.reducerPath]: companyApi.reducer,
  },
  devTools: true,
  middleware: (getDefaultMiddleware) => getDefaultMiddleware()
    .concat(userApi.middleware)
    .concat(authApi.middleware)
    .concat(roleApi.middleware)
    .concat(companyApi.middleware)
});


setupListeners(store.dispatch);

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch