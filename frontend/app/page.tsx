import Link from "next/link";

export default function Home() {
  return (
    <>
    <h1>Hi!</h1>
    <Link href={"/reception"}>受付</Link>
    </>
  );
}
