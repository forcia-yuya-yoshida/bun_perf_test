import { Elysia } from "elysia";

const app = new Elysia();
app.get("/", () => "Hello World!");

app.get("/calc", () => {
  let x = 0;
  for (let i = 1; i < 100000000; i++) {
    x++;
  }
  return x;
});

app.listen(3000);
