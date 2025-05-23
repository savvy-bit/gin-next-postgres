package repositories

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/savvy-bit/gin-react-postgres/dto"
	"github.com/savvy-bit/gin-react-postgres/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	UploadProfileImage(userID uuid.UUID, profileImage string) (*models.User, error)
	UploadBannerImage(userID uuid.UUID, bannerImage string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, *gorm.DB, error)
	VerifyAuthOtp(userID uuid.UUID) (*models.User, *gorm.DB, error)
	RegenerateAuthOtp(userID uuid.UUID) (*models.User, *gorm.DB, error)
	RegenerateAuthTokens(userID uuid.UUID) (*models.User, *gorm.DB, error)
	LoginUser(loginReq dto.UserLoginRequest) (*models.User, *gorm.DB, error)
	LogoutUser(userID uuid.UUID) error
	GetUserByID(userID uuid.UUID) (*models.User, error)
	UpdateUser(userID uuid.UUID, updateReq dto.UserUpdateRequest) (*models.User, error)
	DeleteUser(userID uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(user *models.User) (*models.User, *gorm.DB, error) {
	if err := user.BeforeCreate(u.db); err != nil {
		return nil, nil, err
	}
	var existingUser models.User
	if err := u.db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := u.db.Create(&user).Error; err != nil {
				return nil, nil, err
			}
			return user, u.db, nil
		}
		return nil, nil, err
	}

	return nil, nil, fmt.Errorf("user with email already exists")
}

func (u *userRepository) VerifyAuthOtp(userID uuid.UUID) (*models.User, *gorm.DB, error) {
	var user models.User
	if err := u.db.Model(&user).Where("user_id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, fmt.Errorf("user not found")
		}
		return nil, nil, err
	}
	return &user, u.db, nil
}

func (u *userRepository) RegenerateAuthOtp(userID uuid.UUID) (*models.User, *gorm.DB, error) {
	var user models.User
	if err := u.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, fmt.Errorf("user not found")
		}
		return nil, nil, err
	}
	return &user, u.db, nil
}

func (u *userRepository) LoginUser(loginReq dto.UserLoginRequest) (*models.User, *gorm.DB, error) {
	var user models.User
	if err := u.db.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, fmt.Errorf("user not found")
		}
		return nil, nil, err
	}
	return &user, u.db, nil
}

func (u *userRepository) RegenerateAuthTokens(userID uuid.UUID) (*models.User, *gorm.DB, error) {
	var user models.User
	if err := u.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, fmt.Errorf("user not found")
		}
		return nil, nil, err
	}
	return &user, u.db, nil
}

func (u *userRepository) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, "user_id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) LogoutUser(userID uuid.UUID) error {
	var user models.User
	if err := u.db.Model(&user).Where("user_id = ?", userID).Updates(map[string]interface{}{
		"refresh_token":             nil,
		"refresh_token_expiry_time": nil,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(userID uuid.UUID) error {
	var user models.User
	if err := u.db.Delete(&user, "user_id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateUser(userID uuid.UUID, updateReq dto.UserUpdateRequest) (*models.User, error) {
	var user models.User
	if err := u.db.Model(&user).Where("user_id = ?", userID).Updates(map[string]any{
		"full_name": updateReq.FullName,
		"username":  updateReq.Username,
		"gender":    updateReq.Gender,
	}).Error; err != nil {
		return nil, err
	}
	if err := u.db.First(&user, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) UploadBannerImage(userID uuid.UUID, bannerImage string) (*models.User, error) {
	var user models.User
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{}).Where("user_id = ?", userID).Update("banner_image", bannerImage).Error; err != nil {
			return err
		}
		if err := tx.First(&user, "user_id = ?", userID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *userRepository) UploadProfileImage(userID uuid.UUID, profileImage string) (*models.User, error) {
	var user models.User

	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{}).Where("user_id = ?", userID).Update("profile_image", profileImage).Error; err != nil {
			return err
		}
		if err := tx.First(&user, "user_id = ?", userID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &user, nil
}
