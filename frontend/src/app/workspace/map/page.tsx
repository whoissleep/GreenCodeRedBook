"use client"
import { YMaps, Map, Polygon } from '@pbe/react-yandex-maps';
import Image from "next/image";
import Link from "next/link";
import { Form, useForm, SubmitHandler} from "react-hook-form";
import { useRouter } from "next/navigation";
import logo from "../../../../public/img/favicon.png"

interface IFormStateCadanom{
  number_reestr: string;
}

export default function map() {
  const {register, handleSubmit} = useForm<IFormStateCadanom>();
  const router = useRouter();

  const onSubmit: SubmitHandler<IFormStateCadanom>= async (data) => {
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
      <main className="grid justify-items-center">
      <div className='m-20 flex items-start gap-40'>
        <div className='bg-white rounded-md'>
          <form onSubmit={handleSubmit(onSubmit)} className="py-3 px-10 justify-self-center grid space-y-4">
            <a>Номер кадастрового реестра</a>
            <div className='grid grid-cols-2 gap-4'>
              <input placeholder='Номер реестра' type='text' {...register ('number_reestr')}/>
              <button>Поиск</button>
            </div>
          </form>
        </div>
            <YMaps>
                <Map defaultState={{ center: [55.770000, 37.427884], zoom: 11}} height={500} width={700}>
                <Polygon
      geometry={[
        [
          [55.75, 37.8],
          [55.8, 37.9],
          [55.75, 38.0],
          [55.7, 38.0],
          [55.7, 37.8],
        ],
        [
          [55.75, 37.82],
          [55.75, 37.98],
          [55.65, 37.9],
        ],
      ]}
      options={{
        fillColor: "#00FF00",
        strokeColor: "#0000FF",
        opacity: 0.5,
        strokeWidth: 5,
        strokeStyle: "shortdash",
      }}
    />
    </Map>
            </YMaps>
        </div>
      </main>
      <footer className="bg-[#17A34A] bg-cover h-14 p-10">
        <a className="text-left justify-items-start text-white">© 2024 Москва. ЭкоКарта. Все права защищены.</a>
      </footer>
    </div>
  );
}


