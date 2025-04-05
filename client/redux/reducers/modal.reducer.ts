import { createSlice, PayloadAction } from '@reduxjs/toolkit';

import { ModalState } from '@/types/reducer-state.type';

export const initialCommonState: ModalState = {
  isShowSettingModal: false,
};

export const queryDataSlice = createSlice({
  name: 'auth',
  initialState: initialCommonState,
  reducers: {
    changeIsShowSettingModal: (state: ModalState, action: PayloadAction<boolean>) => {
      state.isShowSettingModal = action.payload;
    },
  },
});

export const {
  changeIsShowSettingModal
} = queryDataSlice.actions;
export default queryDataSlice.reducer;
