/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// [START slides_create_presentation]
package snippets

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/slides/v1"
	"log"
)

func createPresentation() string {
	ctx := context.Background()
	// Uses env GOOGLE_APPLICATION_CREDENTIALS
	client, err := google.DefaultClient(ctx,
		drive.DriveScope,
		slides.PresentationsScope,
		sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Error creating Google client: %v", err)
	}
	slidesService, err := slides.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Slides client: %v", err)
	}
	// Create a presentation and request a PresentationId.
	p := &slides.Presentation{
		Title: "Title",
	}
	presentation, err := slidesService.Presentations.Create(p).Fields(
		"presentationId",
	).Do()
	if err != nil {
		log.Fatalf("Unable to create presentation. %v", err)
	}
	fmt.Printf("Created presentation with ID: %s", presentation.PresentationId)
	return presentation.PresentationId
}

// [END slides_create_presentation]
