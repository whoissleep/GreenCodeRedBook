import Image from "next/image";
import Link from "next/link";
import birds from "../../public/img/headBird.png"

export default function Home() {
  return (
    <div>
      <header>
        <Image src={birds} alt="birds" width={1980}></Image>
      </header>
      <main className="grid justify-items-center">
        <div className="grid justify-items-left">
          <label className="text-8xl text-left bg-[#23C55E] px-5">ЭКОКАРТА</label>
          <a className="w-96 text-2xl text-left"> — это вклад в сохранение биоразнообразия и устойчивое развитие территорий.</a>
        </div>
        <Link href="/workspace/loginPage">
        <button className="mt-20 px-16 py-2 rounded-lg text-xl">Вход</button>
        </Link>
        <Link className="text-white" href="./workspace/map">Перейти на карту</Link>
      </main>
    </div>
  );
}
