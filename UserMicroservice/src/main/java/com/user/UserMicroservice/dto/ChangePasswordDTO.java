package com.user.UserMicroservice.dto;

import com.user.UserMicroservice.validation.ValidPassword;

public class ChangePasswordDTO {

    private String oldPassword;
    @ValidPassword
    private String newPassword;

    public String getOldPassword() { return oldPassword; }

    public void setOldPassword(String oldPassword) {
        this.oldPassword = oldPassword;
    }

    public String getNewPassword() {
        return newPassword;
    }

    public void setNewPassword(String newPassword) {
        this.newPassword = newPassword;
    }
}
