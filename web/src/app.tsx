import { Plus } from "lucide-react";
import { Button } from "./components/ui/button";
import { Dialog, DialogTrigger } from "./components/ui/dialog";
import { CreateGoal } from "./components/create-goal";

export function App() {
  return (
    <Dialog>
      <div className="h-screen flex flex-col items-center justify-center gap-8">
        <p className="text-zinc-300 leading-relaxed max-w-80 text-center">
          Você ainda não cadastrou nenhuma meta, que tal cadastrar uma agora
          mesmo?
        </p>

        <DialogTrigger asChild>
          <Button type="button">
            <Plus className="size-4" />
            Cadastrar meta
          </Button>
        </DialogTrigger>
      </div>

      <CreateGoal />
    </Dialog>
  );
}
