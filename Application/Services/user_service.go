package Services

//
//import (
//	"errors"
//	"go/foodappbe/Application/Services/DTO"
//	domain "go/foodappbe/Domain/Models"
//	"time"
//)
//
//type UserService interface {
//	UpdateUserInfo(input DTO.UpdateUserDto)
//	GetById(id int) (*DTO.UserResponseDto, error)
//	AddAddressForUser(input DTO.CreateAddressDto) error
//	GetAddressForUserPaging(input DTO.AddressFilterDto) (*DTO.PageResultDto, error)
//	UpdateAddressForUser(input DTO.UpdateAddressDto) error
//}
//
//func (s *userServiceImpl) UpdateUserInfo(input DTO.UpdateUserDto) error {
//	user := &domain.User{}
//	if err := s.db.First(user, input.ID).Error; err != nil {
//		return errors.New("user not found")
//	}
//	user.FirstName = input.FirstName
//	user.LastName = input.LastName
//	user.PhoneNumber = input.PhoneNumber
//	return s.db.Save(user).Error
//}
//
//func (s *userServiceImpl) GetById(id int) (*DTO.UserResponseDto, error) {
//	user := &domain.User{}
//	if err := s.db.First(user, id).Error; err != nil {
//		return nil, errors.New("user not found")
//	}
//	return &DTO.UserResponseDto{
//		FirstName:   user.FirstName,
//		LastName:    user.LastName,
//		Email:       user.Email,
//		PhoneNumber: user.PhoneNumber,
//	}, nil
//}
//
//func (s *userServiceImpl) AddAddressForUser(input DTO.CreateAddressDto) error {
//	address := &domain.UserAddress{
//		AddressType:   input.AddressType,
//		CreatedAt:     time.Now(),
//		CreatedBy:     input.UserID,
//		Province:      input.Province,
//		UserID:        input.UserID,
//		DetailAddress: input.DetailAddress,
//		District:      input.District,
//		Notes:         input.Notes,
//		StreetAddress: input.StreetAddress,
//		Ward:          input.Ward,
//	}
//	return s.db.Create(address).Error
//}
//
//func (s *userServiceImpl) GetAddressForUserPaging(input DTO.AddressFilterDto) (*DTO.PageResultDto, error) {
//	var addresses []domain.UserAddress
//	if err := s.db.Where("user_id = ?", input.UserId).Limit(input.PageSize).Offset((input.PageIndex - 1) * input.PageSize).Find(&addresses).Error; err != nil {
//		return nil, err
//	}
//	totalItems := len(addresses)
//	addressDtos := make([]DTO.AddressResponseDto, len(addresses))
//	for i, address := range addresses {
//		addressDtos[i] = DTO.AddressResponseDto{
//			ID:            address.ID,
//			Province:      address.Province,
//			District:      address.District,
//			Ward:          address.Ward,
//			StreetAddress: address.StreetAddress,
//			DetailAddress: address.DetailAddress,
//			Notes:         address.Notes,
//			AddressType:   address.AddressType,
//			CreatedAt:     address.CreatedAt,
//			CreatedBy:     address.CreatedBy,
//			UpdatedAt:     address.UpdatedAt,
//			UpdateBy:      address.UpdateBy,
//		}
//	}
//	return &DTO.PageResultDto{
//		Items:      addressDtos,
//		TotalItems: totalItems,
//	}, nil
//}
//
//func (s *userServiceImpl) UpdateAddressForUser(input DTO.UpdateAddressDto) error {
//	address := &domain.UserAddress{}
//	if err := s.db.First(address, input.Id).Error; err != nil {
//		return errors.New("address not found")
//	}
//	address.UserId = input.UserId
//	address.Province = input.Province
//	address.District = input.District
//	address.Ward = input.Ward
//	address.StreetAddress = input.StreetAddress
//	address.DetailAddress = input.DetailAddress
//	address.Notes = input.Notes
//	address.AddressType = input.AddressType
//	address.UpdatedAt = time.Now()
//	address.UpdateBy = input.UserId
//	return s.db.Save(address).Error
//}
//
//type userServiceImpl struct {
//	db *db.FoodyAppContext
//}
//
//func NewUserService(db *db.FoodyAppContext) UserService {
//	return &userServiceImpl{db: db}
//}