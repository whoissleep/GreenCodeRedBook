"use client"
import Image from "next/image";
import Link from "next/link";
import { Form, useForm, SubmitHandler} from "react-hook-form";
import { useRouter } from "next/navigation";
import logo from "../../../../public/img/favicon.png"

interface IFormStateSignIn{
  name: string;
  password:string;
  email:string;
  phone:string;
}

export default function SigninPage() {
  const {register, handleSubmit} = useForm<IFormStateSignIn>();
  const router = useRouter();

  const onSubmit: SubmitHandler<IFormStateSignIn>= async (data) => {
    console.log("отправляю данные")
      try {
        const response = await fetch('http://localhost:3001/requests', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        })
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const result = await response.json();
      console.log('Success:', result);
      } catch (error) {
        console.error('Error:', error);
      }
      router.push("http://localhost:3000/workspace/map");
  }
  return (
    <div className="bg-[#CAFFDC] bg-cover">
      <header className="p-5 bg-[#17A34A] grid justify-items-center mb-16">
        <Image src={logo} alt="Красная книга" width={79} className="justify-self-center"/>
      </header>
      <main className="grid justify-items-center">
        <div className="m-10 grid justify-items-center">
          <form onSubmit={handleSubmit(onSubmit)} className="justify-self-center grid space-y-4">
            <h1 className="mb-3 text-center text-2xl">Регистрация</h1>
            <input placeholder='Имя и фамилия пользователя' type='name' {...register('name', { required: true })} />

            <input placeholder='Пароль' type="password"{...register('password', {required: true})} />
            <input placeholder='Почта' type='email' {...register('email', { required: true })} />
            <input placeholder="Номер телефона" type='number'{...register('phone', {required:true})}/>
            <button  type="submit" className="px-5 mt-6 text-xl hover: hover:scale-102 hover:bg-orange-400 duration-300">Зарегистрироваться</button>
          </form>
        </div>
      </main>
      <footer className="bg-[#17A34A] h-14 p-10 mt-48">
        <a className="text-white text-left justify-items-start">© 2024 Москва. ЭкоКарта. Все права защищены.</a>
      </footer>
    </div>
  );
}