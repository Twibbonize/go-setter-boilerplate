package main

import (
	"context"
	"net/http"

	moduleboilerplate "github.com/Twibbonize/go-module-boilerplate-mongodb"
	pb_anymodule "github.com/Twibbonize/go-setter-boilerplate/protos/anymodule/protos"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	anyModuleSetter moduleboilerplate.SetterLib
	anyModuleGetter moduleboilerplate.GetterLib
	pb_anymodule.UnimplementedSetterServer
}

func (s *server) SeedOneByRandId(ctx context.Context, in *pb_anymodule.IngestRequestByRandId) (*pb_anymodule.AnyModuleIngestStatus, error) {

	type Body struct {
		RandId string `json:"randid" binding:"required"`
	}
	_, err := s.anyModuleGetter.GetByRandID(in.RandId)
	if err != nil {
		return &pb_anymodule.AnyModuleIngestStatus{Status: pb_anymodule.EnumStatus_BLANK}, nil
	}
	return &pb_anymodule.AnyModuleIngestStatus{Status: pb_anymodule.EnumStatus_SUCCESS}, nil
}

func (s *server) SeedOneByUUID(ctx context.Context, in *pb_anymodule.IngestRequestByUUID) (*pb_anymodule.AnyModuleIngestStatus, error) {

	type Body struct {
		Uuid string `json:"uuid" binding:"required"`
	}

	_, err := s.anyModuleSetter.FindByUUID(in.Uuid, true)
	if err != nil {
		return &pb_anymodule.AnyModuleIngestStatus{Status: pb_anymodule.EnumStatus_FAIL}, nil
	}

	return &pb_anymodule.AnyModuleIngestStatus{Status: pb_anymodule.EnumStatus_SUCCESS}, nil
}

func (s *server) SeedMany(ctx context.Context, in *pb_anymodule.IngestRequest) (*pb_anymodule.AnyModuleIngestStatus, error) {

	type Body struct {
		RetrievedLengthStr string `json:"retrievedlengthstr" binding:"required"`
		LastObjectIdHex    string `json:"lastobjectidhex" binding:"required"`
		ValidLastUUID      string `json:"validlastuuid" binding:"required"`
		CampaignUUID       string `json:"campaignuuid" binding:"required"`
	}

	// Parse retrievedLengthStr as int

	errFind := s.anyModuleSetter.SeedLinked(in.RetrievedLength, in.LastObjectIdHex, in.ValidLastUUID, in.AnyUUID)
	if errFind != nil {
		return &pb_anymodule.AnyModuleIngestStatus{Status: pb_anymodule.EnumStatus_FAIL}, nil
	}

	return &pb_anymodule.AnyModuleIngestStatus{Status: pb_anymodule.EnumStatus_SUCCESS}, nil
}

func (s *server) DeleteManyByAnyUUID(ctx context.Context, in *pb_anymodule.DeleteAllRequest) (*pb_anymodule.AnyModuleIngestStatus, error) {

	type Body struct {
		Uuid string `json:"uuid" binding:"required"`
	}

	err := s.anyModuleSetter.DeleteManyByAnyUUID(in.Uuid)
	if err != nil {
		return &pb_anymodule.AnyModuleIngestStatus{Status: pb_anymodule.EnumStatus_FAIL}, nil
	}

	return &pb_anymodule.AnyModuleIngestStatus{Status: pb_anymodule.EnumStatus_SUCCESS}, nil
}

func Create(c *fiber.Ctx, anyModuleSetter *moduleboilerplate.SetterLib) error {

	type CreateRequest struct {
		AnyUUID string `json:"anyuuid" binding:"required"`
	}
	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	newEntity, err := moduleboilerplate.Init()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to initialize entity"})
	}
	newEntity.AnyUUID = req.AnyUUID

	if setError := anyModuleSetter.Create(newEntity); setError != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": setError.Message})
	}

	return c.Status(http.StatusCreated).JSON(newEntity)
}

func Update(c *fiber.Ctx, anyModuleSetter *moduleboilerplate.SetterLib) error {

	type UpdateRequest struct {
		UUID   string `json:"uuid" binding:"required"`
		RandID string `json:"randid"`
	}
	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	entity, setError := anyModuleSetter.FindByUUID(req.UUID, true)
	if setError != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": setError.Message})
	}

	if req.RandID != "" {
		entity.RandID = req.RandID
	}

	if setError := anyModuleSetter.Update(entity); setError != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": setError.Message})
	}

	return c.Status(http.StatusOK).JSON(entity)
}

func Delete(c *fiber.Ctx, anyModuleSetter *moduleboilerplate.SetterLib) error {

	type DeleteRequest struct {
		UUID string `json:"uuid" binding:"required"`
	}
	var req DeleteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	entity, setError := anyModuleSetter.FindByUUID(req.UUID, true)
	if setError != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": setError.Message})
	}

	if setError := anyModuleSetter.Delete(entity); setError != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": setError.Message})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Entity deleted successfully"})
}