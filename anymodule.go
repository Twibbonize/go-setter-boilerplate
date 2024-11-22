package main

import (
	"context"

	moduleboilerplate "github.com/Twibbonize/go-module-boilerplate-mongodb"
	pb_anymodule "github.com/Twibbonize/go-setter-boilerplate/protos/anymodule/protos"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	anyModuleSetter moduleboilerplate.SetterLib
	anyModuleGetter moduleboilerplate.GetterLib
	pb_anymodule.UnimplementedSetterServer
}

func (s *server) SeedOneByRandId(ctx context.Context, in *pb_anymodule.IngestRequestByRandId) (*pb_anymodule.IngestStatus, error) {

	type Body struct {
		RandId string `json:"randid" binding:"required"`
	}
	_, err := s.anyModuleGetter.GetByRandID(in.RandId)
	if err != nil {
		return &pb_anymodule.IngestStatus{Status: &pb_anymodule.IngestStatus{}}, nil
	}
	return &pb_anymodule.IngestStatus{Status: &pb_anymodule.IngestStatus{}}, nil
}

func (s *server) SeedOneByUUID(ctx context.Context, in *pb_anymodule.IngestRequestByUUID) (*pb_anymodule.IngestStatus, error) {

	type Body struct {
		Uuid string `json:"uuid" binding:"required"`
	}

	err := s.anyModuleSetter.SeedByUUID(in.Uuid)

	if err != nil {
		return &pb_anymodule.IngestStatus{Status: &pb_anymodule.IngestStatus{}}, nil
	}

	return &pb_anymodule.IngestStatus{Status: &pb_anymodule.IngestStatus{}}, nil
}

func (s *server) SeedMany(ctx context.Context, in *pb_anymodule.IngestRequest) (*pb_anymodule.IngestStatus, error) {

	type Body struct {
		RetrievedLengthStr string `json:"retrievedlengthstr" binding:"required"`
		LastObjectIdHex    string `json:"lastobjectidhex" binding:"required"`
		ValidLastUUID      string `json:"validlastuuid" binding:"required"`
		CampaignUUID       string `json:"campaignuuid" binding:"required"`
	}

	// Parse retrievedLengthStr as int

	errFind := s.anyModuleSetter.SeedLinked(in.RetrievedLength, in.LastObjectIdHex, in.ValidLastUUID, in.AnyUUID)
	if errFind != nil {
		return &pb_anymodule.IngestStatus{Status: &pb_anymodule.IngestStatus{}}, nil
	}

	return &pb_anymodule.IngestStatus{Status: &pb_anymodule.IngestStatus{}}, nil
}

func (s *server) DeleteManyByAnyUUID(ctx context.Context, in *pb_anymodule.DeleteAllRequest) (*pb_anymodule.IngestStatus, error) {

	type Body struct {
		Uuid string `json:"uuid" binding:"required"`
	}

	err := s.anyModuleSetter.DeleteManyByAnyUUID(in.Uuid)
	if err != nil {
		return &pb_anymodule.IngestStatus{Status: &pb_anymodule.IngestStatus{}}, nil
	}

	return &pb_anymodule.IngestStatus{Status: &pb_anymodule.IngestStatus{}}, nil
}

func Create(c *fiber.Ctx, anyModuleSetter moduleboilerplate.SetterLib) error {
	responseData := "test"
	return c.Status(fiber.StatusOK).JSON(responseData)

}

func Update(c *fiber.Ctx, anyModuleSetter moduleboilerplate.SetterLib) error {
	responseData := "test"
	return c.Status(fiber.StatusOK).JSON(responseData)

}

func Delete(c *fiber.Ctx, anyModuleSetter moduleboilerplate.SetterLib) error {
	responseData := "test"
	return c.Status(fiber.StatusOK).JSON(responseData)

}
