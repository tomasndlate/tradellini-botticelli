import { startAction, stopAction } from "@/actions/print.action";

export default function Home() {
  return (
    <div className="">
      Test Page
      Start
      <form action={startAction}>
        <textarea name="name"></textarea>
        <button type="submit">Start</button>
      </form>

      Stop
      <form action={stopAction}>
        <textarea name="name"></textarea>
        <button type="submit">Start</button>
      </form>
    </div>
  );
}
