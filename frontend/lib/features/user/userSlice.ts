// src/lib/features/user/userSlice.ts
import { createSlice, PayloadAction } from '@reduxjs/toolkit';

// 1. Definimos o formato (tipo) do estado do usuário
interface UserState {
  name: string;
  email: string;
  picture?: string;
  isLoggedIn: boolean;
}

// 2. Definimos os valores iniciais (quando o app acaba de abrir)
const initialState: UserState = {
  name: '',
  email: '',
  picture: '',
  isLoggedIn: false,
};

// 3. Criamos o Slice com as funções que modificam o estado (reducers)
export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    // Ação para salvar o usuário completo ao fazer login
    loginUser: (state, action: PayloadAction<{ name: string; email: string; picture?: string }>) => {
      state.name = action.payload.name
      state.email = action.payload.email
      state.picture = action.payload.picture
      state.isLoggedIn = true
    },
    // Ação para apenas atualizar o nome se necessário
    updateName: (state, action: PayloadAction<string>) => {
      state.name = action.payload;
    },
    // Ação para deslogar e limpar os dados
    logoutUser: (state) => {
      state.name = ''
      state.email = ''
      state.isLoggedIn = false
    },
  },
});

// 4. Exportamos as ações para usar nos componentes com o dispatch
export const { loginUser, updateName, logoutUser } = userSlice.actions;

// 5. Exportamos o reducer para colocar lá na nossa Store
export default userSlice.reducer;