import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useLoginUserMutation } from "@/redux/services/auth";
import { Loader } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import React from "react";
import { Form, redirect, useNavigate } from "react-router-dom";


interface FormElements extends HTMLInputElement {
  username: HTMLInputElement;
  password: HTMLInputElement;
}


export default function Login() {
  const [loginUser, { isLoading }] = useLoginUserMutation();
  const navigate = useNavigate()


  const handleSubmit = (event: React.SyntheticEvent): void => {
    event.preventDefault();

    const target = event.target as FormElements;
    const username = target["username"]?.value;
    const password = target["password"]?.value;

    loginUser({ username, password }).unwrap().then((data) => {
      console.log({ data });
      if (!data) throw new Error('Login failed')
      if (data.env.toLowerCase() === "development") document.cookie = `Authorization=${data.token}`;
      navigate('/dashboard', { replace: true })
    }).catch((err: string) => {
      console.log({ err });
      notifications.show({ message: err, color: "red", withBorder: true });
    });
  }


  return (
    <>
      {isLoading &&
        <div className=" h-screen w-screen z-50 absolute flex justify-center items-center backdrop-blur-sm">
          <Loader />
        </div>
      }
      <Form onSubmit={(event: React.SyntheticEvent): void => handleSubmit(event)}
        method="post" action="/login" className="grid w-full max-w-sm items-center gap-1.5 mx-auto min-h-screen content-center">
        <Label>Username</Label>
        <Input className="max-w-sm" required name="username" />
        <Label>Password</Label>
        <Input className="max-w-sm" type="password" required name="password" />
        <Button type="submit">Login</Button>
      </Form>
    </>
  )
}

// export async function LoginRouteAction({ request }: { request: Request }) {
//   const data = await request.formData();
//   const username = data.get('username');
//   const password = data.get('password');
//   console.log({ username, password });
//   const [loginUser] = useLoginUserMutation();
//   const login =await loginUser({username, password});
//   console.log({login})
// }