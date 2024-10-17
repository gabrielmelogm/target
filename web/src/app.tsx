import { Plus } from "lucide-react";
import { Button } from "./components/ui/button";

export function App() {
  return (
    <div className="h-screen flex flex-col items-center justify-center gap-8">
      <p className="text-zinc-300 leading-relaxed max-w-80 text-center">
        Você ainda não cadastrou nenhuma meta, que tal cadastrar uma agora
        mesmo?
      </p>

      <Button type="button" variant="primary">
        <Plus className="size-4" />
        Cadastrar meta
      </Button>
    </div>
  );
}
