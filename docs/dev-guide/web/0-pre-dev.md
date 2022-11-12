# Before Development

So this will be a somewhat so-called Vue.ts project on the front-end side, and a Go project on the back-end side. The front-end side will be developed with Vue.js and TypeScript, and the back-end side will be developed with Go and also Kratos, if necessary.

## Setup

> 尤大：“`<script setup>` + TS + Volar = 真香”

```bash
cd web
npm init vue@latest bookkeepingo

...

cd bookkeepingo
npm install
npm run lint
npm run dev
```

And then there prints

```bash
> bookkeepingo@0.0.0 dev
> vite


  VITE v3.2.2  ready in 892 ms

  ➜  Local:   http://localhost:5173/
  ➜  Network: use --host to expose
```

And we can see the "You did it! You’ve successfully created a project with Vite + Vue 3. What's next?" page at `http://localhost:5173/`.
