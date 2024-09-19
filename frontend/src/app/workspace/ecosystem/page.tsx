"use client"
import Image from "next/image";
import Link from "next/link";
import { Form, useForm, SubmitHandler} from "react-hook-form";
import { useRouter } from "next/navigation";
import logo from "../../../../public/img/favicon.png"

interface IFormStateLogin{
  name: string;
  password:string;
  email:string;
  phone:string;
  save_login:boolean;
}

export default function LoginPage() {
  const {register, handleSubmit} = useForm<IFormStateLogin>();
  const router = useRouter();

  const onSubmit: SubmitHandler<IFormStateLogin>= async (data) => {
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
  }
  return (
    <div className="bg-[#CAFFDC] bg-cover">
      <header className="p-3 bg-[#17A34A] grid justify-items-center">
        <Image src={logo} alt="Красная книга" width={79} className="justify-self-center"/>
      </header>
      <nav className="grid grid-cols-2">
        <div className="grid grid-cols-3">
          <Link href="http://localhost:3000/workspace/map">Карта</Link>
          <Link href="http://localhost:3000/workspace/ecosystem">Экосистема</Link>
          <Link href="http://localhost:3000/workspace/notes">Заметки</Link>
        </div>
        <div className="grid justify-items-end">
          <Link href="http://localhost:3000/workspace/UserID">
            <button>Личный кабинет</button>
          </Link>
        </div>
        
      </nav>
      <main className="grid justify-items-center p-96">
        Здесь будет отображение списка красной книги
      </main>
      <footer className="bg-[#17A34A] bg-cover h-14 p-10 mt-40">
        <a className="text-left justify-items-start text-white">© 2024 Москва. ЭкоКарта. Все права защищены.</a>
      </footer>
    </div>
  );
}
