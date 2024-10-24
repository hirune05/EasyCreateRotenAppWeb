import SelectEvent from "@/features/routes/login/components/SelectEvent";
import InputUser from "@/features/routes/login/components/InputUser";
import { Button } from "@/components/ui/button";

 function Login() {
    return (
      <>
  <SelectEvent/>
  <InputUser/>
  <Button className='bg-green-400 text-white py-2 px-4 rounded'>
          会計
        </Button>
      </>
    )
  }

  export default Login